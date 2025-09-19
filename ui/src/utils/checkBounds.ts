function isInNoida(lat: number, lon: number) {
  return (
    lat >= 28.4945 && lat <= 28.6200 &&
    lon >= 77.2830 && lon <= 77.4300
  );
}

export { isInNoida }