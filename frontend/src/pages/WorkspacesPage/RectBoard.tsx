import React from 'react';

const RectBoard: React.FC<{
  name?: string,
}> = ({
  name = "Notes",
}) => {
    return (
      <div className="bg-yellow-600 w-48 h-24 mb-5 mr-5 p-2 rounded">
        <p className="text-base text-white font-bold leading-5">{name}</p>
      </div>
    );
  }

export default RectBoard;
