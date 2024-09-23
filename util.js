import { exists, writeFile } from 'node:fs/promises';
import path from 'path';

let svgImage = '';

export async function makeProgressCircle(_progress) {
  let progress = (Math.max(0, Math.min(100, _progress))).toFixed(1);

  if (isNaN(progress) || !progress)
    progress = 0;

  const filename = path.resolve('/', 'tmp', '.mpris-timer', `_.${progress}.svg`);

  if (await exists(filename)) {
    return filename;
  }

  const width = 256;
  const height = 256;
  const padding = 16;
  const centerX = width / 2;
  const centerY = height / 2;
  const strokeWidth = 32;
  const radius = width / 2 - strokeWidth - padding;
  const baseWidth = Math.floor(strokeWidth * 0.25);

  const red = Math.min(255, Math.floor(255 * (progress / 100)));
  const green = Math.min(255, Math.floor(255 * ((100 - progress) / 100)));

  svgImage = `
    <svg width="${width}" height="${height}">
      <circle 
        cx="${centerX}" 
        cy="${centerY}" 
        r="${radius}" 
        fill="none"
        stroke="#535353"
        stroke-width="${baseWidth}"
      />
      <circle 
        cx="${centerX}" 
        cy="${centerY}" 
        r="${radius}" 
        fill="none"
        stroke="rgb(${red},${green},64)"
        stroke-width="${strokeWidth}"
        stroke-dasharray="${2 * Math.PI * radius}"
        stroke-dashoffset="${2 * Math.PI * radius * (1 - progress / 100)}"
        transform="rotate(-90 ${centerX} ${centerY})"
      />
    </svg>
  `;

  await writeFile(filename, Buffer.from(svgImage.trim()));

  return filename;
}

export const formatMilliseconds = ms => {
  const totalSeconds = Math.floor(ms / 1000);
  const hours = Math.floor(totalSeconds / 3600);
  const minutes = Math.floor((totalSeconds % 3600) / 60);
  const seconds = totalSeconds % 60;

  if (hours > 0) {
    return [
      hours.toString().padStart(2, '0'),
      minutes.toString().padStart(2, '0'),
      seconds.toString().padStart(2, '0')
    ].join(':');
  } else {
    return [
      minutes.toString().padStart(2, '0'),
      seconds.toString().padStart(2, '0')
    ].join(':');
  }
}
