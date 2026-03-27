import { FaCarSide } from 'react-icons/fa';

const Cab = ({
  onClick,
  selected,
}: {
  onClick?: () => void;
  selected?: boolean;
}) => (
  <button
    aria-label='Cab'
    className={`p-4 rounded-md transition hover:bg-surface-selected cursor-pointer flex items-center justify-center${selected ? ' bg-surface-selected' : ''}`}
    onClick={onClick}
  >
    <FaCarSide
      size={24}
      className={`${selected ? 'text-cab' : 'text-cab-muted'} hover:text-cab`}
    />
  </button>
);

export default Cab;
