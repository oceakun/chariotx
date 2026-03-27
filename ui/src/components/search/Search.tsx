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
      className={`p-4 rounded-md transition hover:bg-surface-selected cursor-pointer ${selected ? 'bg-surface-selected' : ''}`}
    >
      <FiSearch
        size={24}
        className={`${selected ? 'text-search' : 'text-search-muted'} hover:text-search`}
      />
    </button>
  );
};

export default Search;
