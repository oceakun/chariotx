'use client';

import React from 'react';

function LoadingScreen() {
  return (
    <div className='absolute inset-0 z-20 flex flex-col items-center justify-center bg-surface'>
      <div className='flex flex-col items-center gap-8'>
        {/* Logo / wordmark */}
        <div className='flex flex-col items-center gap-2'>
          <span className='text-4xl font-black tracking-widest text-text uppercase'>
            Chariot<span className='text-search'>X</span>
          </span>
          <span className='text-xs tracking-[0.3em] text-text-muted uppercase'>
            Loading your map
          </span>
        </div>

        {/* Animated bar */}
        <div className='w-48 h-[2px] bg-surface-selected rounded-full overflow-hidden'>
          <div className='h-full w-1/2 bg-search rounded-full animate-loading-bar' />
        </div>
      </div>
    </div>
  );
}

export default LoadingScreen;
