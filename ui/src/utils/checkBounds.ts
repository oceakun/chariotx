function isInNoida(lat: number, lon: number) {
  return lat >= 28.4945 && lat <= 28.62 && lon >= 77.283 && lon <= 77.43;
}

export { isInNoida };
