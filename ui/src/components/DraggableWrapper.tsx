"use client";
import Draggable from "react-draggable";
import React, { useRef } from "react";

interface DraggableWrapperProps {
  children: React.ReactNode;
}

const DraggableWrapper: React.FC<DraggableWrapperProps> = ({ children }) => {
  const nodeRef = useRef<HTMLDivElement>(null);

  return (
    <Draggable handle=".draggable-handle" nodeRef={nodeRef}>
      <div ref={nodeRef}>{children}</div>
    </Draggable>
  );
};

export default DraggableWrapper;