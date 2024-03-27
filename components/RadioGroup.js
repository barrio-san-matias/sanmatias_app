// components/RadioGroup.js

const RadioGroup = ({ selectedOption, onOptionChange }) => {
  const handleOptionChange = (event) => {
    const value = event.target.value;
    onOptionChange(value); // Callback to update state in the parent
  };

  return (
    <div>
      opciones:
      <label>
        <input
          type="radio"
          value="google"
          checked={selectedOption === "google"}
          onChange={handleOptionChange}
        />
        Google Maps
      </label>
      <label>
        <input
          type="radio"
          value="waze"
          checked={selectedOption === "waze"}
          onChange={handleOptionChange}
        />
        Waze
      </label>
      <label>
        <input
          type="radio"
          value="apple"
          checked={selectedOption === "apple"}
          onChange={handleOptionChange}
        />
        Apple Maps
      </label>
    </div>
  );
};

export default RadioGroup;
