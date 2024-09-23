import path from 'path';
import * as dbus from 'dbus-native';
import * as fs from 'node:fs';
import { makeProgressCircle, formatMilliseconds, FPS } from './util.js';

class MPRISPlayer {
  isPaused = false;
  pausedFor = 0;
  tickTimeout = null;

  constructor(timeSeconds = -1, name = 'Timer') {
    if (!fs.existsSync(path.resolve('/', 'tmp', '.mpris-timer'))) {
      fs.mkdirSync(path.resolve('/', 'tmp', '.mpris-timer'));
    }

    timeSeconds = parseInt(timeSeconds.toString());

    if (!timeSeconds || timeSeconds <= 0 || isNaN(timeSeconds)) {
      console.log('Usage: mpris-timer <seconds> [<title>]');
      process.exit(1);
    }

    const fpsMultiplier = timeSeconds < 90
      ? 1
      : 1 / (timeSeconds / 90);

    this.interval = 1000 / FPS * fpsMultiplier; // fork and do better
    this.name = name;
    this.time = timeSeconds * 1000;
  }

  async init() {
    this.serviceName = `org.mpris.MediaPlayer2.mpris-timer-${Math.random().toString().replace('.', '')}`;
    this.objectPath = '/org/mpris/MediaPlayer2';

    this.metadata = [
      ['mpris:trackid', ['o', '/track/1']],
      ['xesam:title', ['s', this.name]],
      ['xesam:artist', ['as', [ '00:00' ]]],
      ['mpris:artUrl', ['s', `file://${await makeProgressCircle(0)}`]],
    ];

    this.bus = dbus.sessionBus();
    this.playbackStatus = 'Playing';

    this.mprisInterface = {
      name: 'org.mpris.MediaPlayer2',
      methods: {
        Raise: ['', ''],
        Quit: ['', ''],
      },
      properties: {
        Identity: 's',
        DesktopEntry: 's',
      },
    };

    this.playerInterface = {
      name: 'org.mpris.MediaPlayer2.Player',
      methods: {
        PlayPause: ['', ''],
        Stop: ['', ''],
        Next: ['', ''],
        Previous: ['', ''],
      },
      properties: {
        PlaybackStatus: 's',
        Metadata: 'a{sv}',
        Position: 'x',
        CanGoNext: 'b',
        CanGoPrevious: 'b',
        CanPlay: 'b',
        CanPause: 'b',
        CanSeek: 'b',
        CanControl: 'b',
      },
    };

    this.bus.requestName(this.serviceName, 0x4, (err) => {
      this.startTime = Date.now();

      if (err) {
        console.log(`Failed to request service name: ${err}`);
        return;
      }

      this.bus.exportInterface(this, this.objectPath, this.playerInterface);
      this.bus.exportInterface(this, this.objectPath, this.mprisInterface);
      this.tick();
    });
  }

  nextTick() {
    clearTimeout(this.tickTimeout);
    this.tickTimeout = setTimeout(() => this.tick(), 1000 / FPS * this.fpsMultiplier);
  }

  async tick() {
    if (this.isPaused) {
      return this.nextTick();
    }

    const timeDiff = Date.now() - this.startTime - this.pausedFor;
    const timeLeft = this.time - timeDiff;
    const progress = Math.min(1, timeDiff / this.time) * 100;
    const image = await makeProgressCircle(progress);

    if (progress >= 100) {
      process.exit(0);
    }

    const [ , artUrlField ] = this.metadata.find(([ name ]) => name === 'mpris:artUrl');
    artUrlField[1] = `file://${image}`;

    const [ , titleField ] = this.metadata.find(([ name ]) => name === 'xesam:artist');
    titleField[1] = [ formatMilliseconds(timeLeft) ];

    this.emitMetadata();
    this.nextTick();
  }

  emitMetadata() {
    this.bus.sendSignal(
      this.objectPath,
      'org.freedesktop.DBus.Properties',
      'PropertiesChanged',
      'sa{sv}as',
      [
        this.playerInterface.name,
        [
          [ 'Metadata', [ 'a{sv}', this.Metadata ] ],
          [ 'PlaybackStatus', [ 's', this.playbackStatus ] ],
        ],
        [],
      ],
    );
  }

  get Metadata() { return this.metadata }
  get PlaybackStatus() { return this.playbackStatus }
  get Identity() { return 'MPRIS Timer' }
  get DesktopEntry() { return path.resolve(__dirname, 'mpris-timer.desktop') }
  get HasTrackList() { return false }
  get Position() { return 0 }
  get CanGoNext() { return true }
  get CanGoPrevious() { return true }
  get CanPlay() { return true }
  get CanPause() { return true }
  get CanSeek() { return false }
  get CanControl() { return true }

  PlayPause() {
    if (this.isPaused) {
      this.pausedFor += Date.now() - this.pausedAt;
    } else {
      this.pausedAt = Date.now();
    }

    this.isPaused = !this.isPaused;
    this.playbackStatus = this.isPaused ? 'Paused' : 'Playing';
    this.emitMetadata();
  }

  Previous() {
    clearTimeout(this.tickTimeout);

    this.pausedFor = 0;
    this.isPaused = false;
    this.playbackStatus = 'Playing';
    this.startTime = Date.now();
    this.tick();
  }

  Next() { process.exit(1) }
  Stop() { process.exit(1) }
}

const [ ,, time, name ] = process.argv;
const player = new MPRISPlayer(time, name);

player.init();
