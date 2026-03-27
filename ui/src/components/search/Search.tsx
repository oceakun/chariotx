import React from 'react';
import { FiSearch } from 'react-icons/fi';

const Search: React.FC<{ onClick?: () => void; selected?: boolean }> = ({
  onClick,
  selected,
}) => {
  return (
    <button
      aria-label='Open search'
      onClick={onClick}
      className={`p-2 rounded-full transition hover:bg-black cursor-pointer ${selected ? 'bg-black' : ''}`}
    >
      <FiSearch
        size={24}
        className={`${selected ? 'text-yellow-400' : 'text-yellow-200'} hover:text-yellow-400`}
      />
    </button>
  );
};

export default Search;
