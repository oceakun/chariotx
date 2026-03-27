import { FaRegClock } from 'react-icons/fa';

const Estimate = ({
  onClick,
  selected,
}: {
  onClick?: () => void;
  selected?: boolean;
}) => (
  <button
    aria-label='Estimate'
    className={`p-4 rounded-md transition hover:bg-surface-selected cursor-pointer flex items-center justify-center${selected ? ' bg-surface-selected' : ''}`}
    onClick={onClick}
  >
    <FaRegClock
      size={24}
      className={`${selected ? 'text-estimate' : 'text-estimate-muted'} hover:text-estimate`}
    />
  </button>
);

export default Estimate;
