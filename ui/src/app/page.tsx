"use client";

import dynamic from "next/dynamic";
import { useEffect, useState } from "react";
import type { LatLngBoundsExpression } from "leaflet";
import Planner from "@/components/planner/Planner";
import Cab from "@/components/cab/Cab";
import Estimate from "@/components/estimate/Estimate";
import SearchDetails from "@/components/search/Details";
import PlannerDetails from "@/components/planner/Details";
import CabDetails from "@/components/cab/Details";
import EstimateDetails from "@/components/estimate/Details";
import Search from "@/components/search/Search";
import DraggableWrapper from "@/components/DraggableWrapper";
import CloseDetails from "@/components/CloseDetails";
import LoadingScreen from "@/components/LoadingScreen";
import { MdDragIndicator } from "react-icons/md";
import Tracking from "@/components/tracking/Tracking";
import TrackingDetails from "@/components/tracking/Details";

const Map = dynamic(() => import('@/components/map'), { ssr: false });

export default function Page() {
    const [input, setInput] = useState("");
    const [marker, setMarker] = useState<[number, number]>([4.79029, -75.69003]);
    const [bounds, setBounds] = useState<LatLngBoundsExpression | undefined>(undefined);
    const [mounted, setMounted] = useState(false);
    const [selected, setSelected] = useState<string | null>(null);
    const [loading, setLoading] = useState(false);

    useEffect(() => {
        setMounted(true);
    }, []);

    function isInNoida(lat: number, lon: number) {
      return (
        lat >= 28.4945 && lat <= 28.6200 &&
        lon >= 77.2830 && lon <= 77.4300
      );
    }

    async function handleSearch(query: string) {
        if (query.trim()) {
            setLoading(true);
            try {
                // Call Nominatim API
                const res = await fetch(`https://nominatim.openstreetmap.org/search?format=json&q=${encodeURIComponent(query)}`);
                const data = await res.json();
                if (data && data.length > 0) {
                    const lat = parseFloat(data[0].lat);
                    const lon = parseFloat(data[0].lon);
                    if (isInNoida(lat, lon)) {
                      setMarker([lat, lon]);
                      if (data[0].boundingbox) {
                        // boundingbox: [south, north, west, east]
                        const bb = data[0].boundingbox;
                        setBounds([
                          [parseFloat(bb[0]), parseFloat(bb[2])], // SW
                          [parseFloat(bb[1]), parseFloat(bb[3])]  // NE
                        ]);
                      } else {
                        setBounds(undefined);
                      }
                    } else {
                      alert("Can't generate a path for the searched place. Only Noida is supported.");
                    }
                } else {
                    alert("Location not found");
                }
            } finally {
                setLoading(false);
            }
        }
    }

    const isDetailsOpen = Boolean(selected);

    return (
        <div className="flex flex-col h-full w-full">
              <div className="fixed top-0 left-0 right-0 z-[9999]">
    <DraggableWrapper>
      <header className="draggable-handle w-full p-2 bg-transparent flex flex-row items-center justify-center rounded-full h-auto cursor-move">
                <div className={`flex flex-row items-start justify-center gap-2 ${isDetailsOpen ? 'bg-black rounded-[15px]' : 'w-full h-full'}`}>
                {/* Feature Options */}
                <div className={`flex border border-gray-700 rounded-[15px] shadow-lg bg-[#303030] ${isDetailsOpen ? 'flex-col items-stretch py-2 px-2 gap-4' : 'flex-row justify-center items-center gap-4 px-10 py-2'} ${isDetailsOpen ? 'h-full' : ''}`}
                  style={isDetailsOpen ? {height: '100%'} : {}}>
                    <Search onClick={() => setSelected('search')} selected={selected === 'search'} />
                    <Planner onClick={() => setSelected('planner')} selected={selected === 'planner'} />
                    <Cab onClick={() => setSelected('cab')} selected={selected === 'cab'} />
                    <Estimate onClick={() => setSelected('estimate')} selected={selected === 'estimate'} />
                    <Tracking onClick={() => setSelected('tracking')} selected={selected === 'tracking'} />
                </div>
                {/* render Details component corresponding to the selected component */}
                {isDetailsOpen && (
                  <div className="flex justify-start max-w-[300px] w-full h-full relative">
                    <div className="absolute top-2 right-2 z-10">
                      <CloseDetails handleClose={() => setSelected(null)} />
                    </div>
                    <div className="w-full">
                      {selected === 'search' && <SearchDetails onSearch={handleSearch} />}
                      {selected === 'planner' && <PlannerDetails />}
                      {selected === 'cab' && <CabDetails />}
                      {selected === 'estimate' && <EstimateDetails />}
                      {selected === 'tracking' && <TrackingDetails />}
                    </div>
                  </div>
                )}
                </div>
            </header>
            
            </DraggableWrapper>
                </div>
            
            <main className="flex-1 flex flex-col items-center justify-center bg-gray-100 h-full">
    <div className="bg-white-700 w-[100%] h-[100vh] relative">
        {loading && <LoadingScreen />}
        {mounted && !loading && <Map posix={marker} bounds={bounds} />}
    </div>
</main>
        </div>
    )
}