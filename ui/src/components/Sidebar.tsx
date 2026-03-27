'use client';

import React from 'react';
import SidebarHeader from '@/components/SidebarHeader';
import FeatureOptions from '@/components/FeatureOptions';

interface SidebarProps {
  open: boolean;
  onClose: () => void;
  selected: string | null;
  onSelect: (feature: string) => void;
  onSearch: (query: string) => Promise<string | null>;
}

export default function Sidebar({
  open,
  onClose,
  selected,
  onSelect,
  onSearch,
}: SidebarProps) {
  return (
    <div
      className={`h-full bg-surface shadow-xl flex flex-col flex-shrink-0 transition-all duration-300 overflow-hidden ${open ? 'w-[360]' : 'w-[0]'}`}
    >
      <SidebarHeader onClose={onClose} />
      <div className='p-4 flex flex-col gap-2'>
        <FeatureOptions
          selected={selected}
          onSelect={onSelect}
          onSearch={onSearch}
        />
      </div>
    </div>
  );
}
