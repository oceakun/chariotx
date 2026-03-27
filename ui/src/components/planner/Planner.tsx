import { FaRoute } from 'react-icons/fa';

const Planner = ({
  onClick,
  selected,
}: {
  onClick?: () => void;
  selected?: boolean;
}) => (
  <button
    aria-label='Planner'
    className={`p-4 rounded-md transition hover:bg-surface-selected cursor-pointer flex items-center justify-center${selected ? ' bg-surface-selected' : ''}`}
    onClick={onClick}
  >
    <FaRoute
      size={24}
      className={`${selected ? 'text-planner' : 'text-planner-muted'} hover:text-planner`}
    />
  </button>
);

export default Planner;
