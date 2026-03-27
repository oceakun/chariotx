'use client';

import dynamic from 'next/dynamic';
import { useEffect, useState } from 'react';
import type { LatLngBoundsExpression } from 'leaflet';
import LoadingScreen from '@/components/LoadingScreen';
import NavButton from '@/components/NavButton';
import Sidebar from '@/components/Sidebar';

const Map = dynamic(() => import('@/components/map'), { ssr: false });

export default function Page() {
  const [input, setInput] = useState('');
  const [marker, setMarker] = useState<[number, number]>([4.79029, -75.69003]);
  const [bounds, setBounds] = useState<LatLngBoundsExpression | undefined>(
    undefined
  );
  const [mounted, setMounted] = useState(false);
  const [selected, setSelected] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [sidebarOpen, setSidebarOpen] = useState(false);

  function toggleSidebar() {
    setSidebarOpen((o) => {
      const next = !o;
      if (next) setSelected('search');
      localStorage.setItem('sidebarOpen', String(next));
      return next;
    });
  }

  useEffect(() => {
    setMounted(true);
    if ((window as any).__SIDEBAR_OPEN__) {
      setSelected('search');
    }
  }, []);

  function isInNoida(lat: number, lon: number) {
    return lat >= 28.4945 && lat <= 28.62 && lon >= 77.283 && lon <= 77.43;
  }

  async function handleSearch(query: string): Promise<string | null> {
    if (!query.trim()) return null;
    setLoading(true);
    try {
      const res = await fetch(
        `https://nominatim.openstreetmap.org/search?format=json&q=${encodeURIComponent(query)}`
      );
      const data = await res.json();
      if (data && data.length > 0) {
        const lat = parseFloat(data[0].lat);
        const lon = parseFloat(data[0].lon);
        if (isInNoida(lat, lon)) {
          setMarker([lat, lon]);
          if (data[0].boundingbox) {
            const bb = data[0].boundingbox;
            setBounds([
              [parseFloat(bb[0]), parseFloat(bb[2])],
              [parseFloat(bb[1]), parseFloat(bb[3])],
            ]);
          } else {
            setBounds(undefined);
          }
          return null;
        } else {
          return 'Location is outside Noida. Only Noida is currently supported.';
        }
      } else {
        return 'Location not found. Try a different search.';
      }
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className='flex h-screen w-screen overflow-hidden'>
      {!sidebarOpen && <NavButton open={sidebarOpen} onClick={toggleSidebar} />}

      <Sidebar
        open={sidebarOpen}
        onClose={toggleSidebar}
        selected={selected}
        onSelect={setSelected}
        onSearch={handleSearch}
      />

      <main className='flex-1 h-full relative'>
        {loading && <LoadingScreen />}
        {mounted && !loading && <Map posix={marker} bounds={bounds} />}
      </main>
    </div>
  );
}
