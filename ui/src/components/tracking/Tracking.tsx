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
      className={`p-2 rounded-full transition hover:bg-black hover:cursor-pointer flex items-center justify-center${selected ? ' bg-black' : ''}`}
      aria-label='Track Package'
    >
      <MdDeliveryDining
        size={24}
        className={`${selected ? 'text-violet-400' : 'text-violet-200'} hover:text-violet-400`}
      />
    </button>
  );
}

export default Tracking;
