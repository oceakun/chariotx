'use client';

import React from 'react';
import { RxHamburgerMenu, RxCross2 } from 'react-icons/rx';

interface NavButtonProps {
  open: boolean;
  onClick: () => void;
}

export default function NavButton({ open, onClick }: NavButtonProps) {
  return (
    <button
      className='fixed top-4 left-4 z-[412] flex items-center gap-2 bg-surface rounded-lg shadow-lg cursor-pointer'
      onClick={onClick}
    >
      <span className='text-text p-2'>
        {open ? <RxCross2 size={20} /> : <RxHamburgerMenu size={20} />}
      </span>
      <span className='text-sm text-text px-3 py-2'>
        ChariotX - Plan, ride, track
      </span>
    </button>
  );
}
