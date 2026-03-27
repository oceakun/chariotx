import React from 'react';
import { IoMdCloseCircle } from 'react-icons/io';

function CloseDetails({ handleClose }: { handleClose: () => void }) {
  return (
    <button
      aria-label='Close details'
      onClick={handleClose}
      className='p-2 rounded-full hover:text-close-hover cursor-pointer transition flex items-center justify-center text-close'
    >
      <IoMdCloseCircle size={24} />
    </button>
  );
}

export default CloseDetails;
