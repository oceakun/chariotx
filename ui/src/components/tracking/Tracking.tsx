import React from 'react';
import { MdDeliveryDining } from 'react-icons/md';

function Tracking({
  selected,
  onClick,
}: {
  selected: boolean;
  onClick: () => void;
}) {
  return (
    <button
      onClick={onClick}
      className={`p-4 rounded-md transition hover:bg-surface-selected cursor-pointer flex items-center justify-center${selected ? ' bg-surface-selected' : ''}`}
      aria-label='Track Package'
    >
      <MdDeliveryDining
        size={24}
        className={`${selected ? 'text-tracking' : 'text-tracking-muted'} hover:text-tracking`}
      />
    </button>
  );
}

export default Tracking;
