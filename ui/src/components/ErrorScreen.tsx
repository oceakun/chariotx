'use client';

import { TbMapOff } from 'react-icons/tb';

interface ErrorScreenProps {
  message: string;
  onDismiss: () => void;
}

function ErrorScreen({ message, onDismiss }: ErrorScreenProps) {
  return (
    <div className='absolute bottom-8 left-1/2 -translate-x-1/2 z-30 animate-error-in'>
      <div className='flex items-center gap-3 bg-surface border border-input-border text-text px-5 py-3 rounded-xl shadow-2xl min-w-[260px] max-w-sm'>
        <TbMapOff size={20} className='text-search shrink-0' />
        <p className='text-sm leading-snug'>{message}</p>
        <button
          onClick={onDismiss}
          className='ml-auto text-text-muted hover:text-text cursor-pointer text-lg leading-none'
          aria-label='Dismiss'
        >
          ×
        </button>
      </div>
    </div>
  );
}

export default ErrorScreen;
