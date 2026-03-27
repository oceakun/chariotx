'use client';

import React from 'react';
import Search from '@/components/search/Search';
import Planner from '@/components/planner/Planner';
import Cab from '@/components/cab/Cab';
import Estimate from '@/components/estimate/Estimate';
import Tracking from '@/components/tracking/Tracking';
import SearchDetails from '@/components/search/Details';
import PlannerDetails from '@/components/planner/Details';
import CabDetails from '@/components/cab/Details';
import EstimateDetails from '@/components/estimate/Details';
import TrackingDetails from '@/components/tracking/Details';

interface FeatureOptionsProps {
  selected: string | null;
  onSelect: (feature: string) => void;
  onSearch: (query: string) => Promise<string | null>;
}

export default function FeatureOptions({
  selected,
  onSelect,
  onSearch,
}: FeatureOptionsProps) {
  return (
    <div className='flex flex-col gap-4 w-[330px]'>
      <div className='flex flex-row justify-between items-center px-0'>
        <Search
          onClick={() => onSelect('search')}
          selected={selected === 'search'}
        />
        <Planner
          onClick={() => onSelect('planner')}
          selected={selected === 'planner'}
        />
        <Cab onClick={() => onSelect('cab')} selected={selected === 'cab'} />
        <Estimate
          onClick={() => onSelect('estimate')}
          selected={selected === 'estimate'}
        />
        <Tracking
          onClick={() => onSelect('tracking')}
          selected={selected === 'tracking'}
        />
      </div>
      {selected && (
        <div className='w-full  mx-auto'>
          {selected === 'search' && <SearchDetails onSearch={onSearch} />}
          {selected === 'planner' && <PlannerDetails />}
          {selected === 'cab' && <CabDetails />}
          {selected === 'estimate' && <EstimateDetails />}
          {selected === 'tracking' && <TrackingDetails />}
        </div>
      )}
    </div>
  );
}
