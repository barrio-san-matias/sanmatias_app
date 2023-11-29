// components/RadioGroup.js
import { useState } from 'react';

const RadioGroup = ({ selectedOption, onOptionChange }) => {
  const handleOptionChange = (event) => {
    const value = event.target.value;
    onOptionChange(value); // Callback to update state in the parent
  };

  return (
    <div>
    <label>
        <input
          type="radio"
          value="google"
          checked={selectedOption === 'google'}
          onChange={handleOptionChange}
        />
        Google Maps
      </label>



      <label>
        <input
          type="radio"
          value="waze"
          checked={selectedOption === 'waze'}
          onChange={handleOptionChange}
        />
        Waze
      </label>
    </div>
  );
};

export default RadioGroup;

