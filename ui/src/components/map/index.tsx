'use client';

import 'leaflet/dist/leaflet.css';
import {
  MapContainer,
  TileLayer,
  Marker,
  Popup,
  useMap,
  ZoomControl,
} from 'react-leaflet';
import { LatLngExpression, LatLngTuple, LatLngBoundsExpression } from 'leaflet';
import { useEffect } from 'react';
import 'leaflet-defaulticon-compatibility/dist/leaflet-defaulticon-compatibility.css';
import 'leaflet-defaulticon-compatibility';
import { useTheme } from '@/context/ThemeContext';

interface MapViewProps {
  posix: LatLngExpression | LatLngTuple;
  bounds?: LatLngBoundsExpression;
  zoom?: number;
}

const defaults = {
  zoom: 19,
};

function MapUpdater({
  posix,
  bounds,
}: {
  posix: LatLngExpression | LatLngTuple;
  bounds?: LatLngBoundsExpression;
}) {
  const map = useMap();
  useEffect(() => {
    if (bounds) {
      map.fitBounds(bounds, { padding: [40, 40] });
    } else {
      map.setView(posix);
    }
  }, [posix, bounds, map]);
  return null;
}

export default function MapView({
  posix,
  bounds,
  zoom = defaults.zoom,
}: MapViewProps) {
  const { dark } = useTheme();

  return (
    <MapContainer
      key={`${JSON.stringify(posix)}-${dark}`}
      center={posix}
      zoom={zoom}
      scrollWheelZoom={false}
      zoomControl={false}
      className={dark ? 'dark-map' : ''}
      style={{ height: '100%', width: '100%' }}
    >
      <ZoomControl position='bottomright' />
      <TileLayer
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        url='https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png'
      />
      <MapUpdater posix={posix} bounds={bounds} />
      <Marker position={posix} draggable={false}>
        <Popup>Here you are!</Popup>
      </Marker>
    </MapContainer>
  );
}
