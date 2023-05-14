import * as React from "react";

export default function AmountSelectorSlider({
  min,
  max,
}: {
  min: number;
  max: number;
}) {
  const [value, setValue] = React.useState(0);

  function print(v: number) {
    console.log(v);
    setValue(v);
  }

  return (
    <div className="flex flex-col">
      <div className="flex flex-row justify-between">
        <div className="text-gray-500 text-xs">0</div>
        <div className="text-gray-500 text-xs">10</div>
      </div>
      <input
        className="w-full"
        step={0.5}
        type="range"
        min={min}
        max={max}
        value={value}
        onChange={(e) => print(e.target.valueAsNumber)}
      />
    </div>
  );
}
