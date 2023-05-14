import * as React from "react";

export default function ToggleableBox({
  title,
  defaultState,
  imageUrl,
  children,
}: {
  title: string;
  defaultState: "open" | "closed";
  imageUrl: string;
  children: React.ReactNode;
}) {
  const [isOpen, setIsOpen] = React.useState(false);
  const toggleBox = () => {
    setIsOpen((isOpen) => !isOpen);
  };

  React.useEffect(() => {
    if (defaultState === "open") {
      setIsOpen(true);
    }
  }, [defaultState]);

  return (
    <div
      className={`relative flex flex-col ${
        isOpen ? "divide-y" : "divide-opacity-0"
      } divide-discord_offline_inner rounded border border-strawberry bg-darkmode_light bg-opacity-80`}
    >
      <button className={`relative h-11 text-white`} onClick={toggleBox}>
        <div
          className={`absolute flex inset-y-0 items-center left-2 transition ${
            isOpen ? "rotate-180" : "rotate-90"
          } duration-300`}
        >
          <svg
            fill="#ffffff"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            width="24"
            height="24"
          >
            <path fill="none" d="M0 0h24v24H0z" />
            <path d="M12 10.828l-4.95 4.95-1.414-1.414L12 8l6.364 6.364-1.414 1.414z" />
          </svg>
        </div>
        <h2 className="flex justify-center items-center absolute font-bold text-gray inset-0">
          {title}
        </h2>
        <div className="flex justify-end absolute inset-0 right-3 items-center">
          <img alt=" " src={imageUrl} width={30} />
        </div>
      </button>
      <div
        className={`overflow-y-hidden duration-300 ease-in-out ${
          isOpen ? "max-h-[120rem]" : "max-h-0"
        }`}
      >
        <div className="m-4">{children}</div>
      </div>
    </div>
  );
}
