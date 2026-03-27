import React from 'react';
import { FiX } from 'react-icons/fi';
import { IoMdCloseCircle } from 'react-icons/io';

function CloseDetails({ handleClose }: { handleClose: () => void }) {
  return (
    <button
      aria-label='Close details'
      onClick={handleClose}
      className='p-2 rounded-full hover:text-red-500 cursor-pointer transition flex items-center justify-center text-red-200'
    >
      <IoMdCloseCircle size={24} />
    </button>
  );
}

export default CloseDetails;
