'use client';

import React from 'react';
import { RxCross2 } from 'react-icons/rx';
import { BsSun, BsMoon } from 'react-icons/bs';
import { useTheme } from '@/context/ThemeContext';

interface SidebarHeaderProps {
  onClose: () => void;
}

export default function SidebarHeader({ onClose }: SidebarHeaderProps) {
  const { dark, toggle } = useTheme();

  return (
    <div className='w-full flex items-center justify-between px-3 py-3 border-b-2 border-surface-selected'>
      <button
        onClick={onClose}
        className='p-1.5 rounded-md text-text-muted hover:text-text hover:bg-surface-selected transition cursor-pointer'
        aria-label='Close sidebar'
      >
        <RxCross2 size={18} />
      </button>

      <button
        onClick={toggle}
        className='flex items-center gap-2 px-3 py-1.5 rounded-full bg-surface-selected text-text-muted hover:text-text transition cursor-pointer text-sm'
        aria-label='Toggle theme'
      >
        {dark ? <BsMoon size={14} /> : <BsSun size={14} />}
        <span>{dark ? 'Dark' : 'Light'}</span>
      </button>
    </div>
  );
}
