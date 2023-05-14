import { NavLink } from "@remix-run/react";
import * as React from "react";

export default function Navbar({
  fixed,
}: {
  /**
   * otherwise default to relative
   */
  fixed: boolean;
}) {
  const [isNavbarOpen, setIsNavbarOpen] = React.useState(true);

  const toggleNavbar = () => {
    setIsNavbarOpen((prev) => !prev);
  };

  return (
    <nav className={`${fixed ? "fixed inset-0" : "relative"} h-fit`}>
      <div
        className={`absolute flex laptop:justify-end ${
          isNavbarOpen
            ? "laptop:w-full inset-x-auto"
            : "laptop:w-0 inset-x-full"
        } laptop:border-b-2 laptop:border-darkmode_light rounded laptop:rounded-none bg-black bg-opacity-70 laptop:bg-opacity-100 laptop:bg-darkmode_dark duration-500 ease-in-out pt-12`}
      >
        <ul className="laptop:h-12 text-center h-fit flex-row laptop:flex items-center space-y-3 text-white font-bold pt-3 laptop:pt-0 bg-opacity-0 laptop:bg-inherit">
          <li className="duration-200 hover:scale-110 hover:underline decoration-strawberry mx-3">
            <NavLink to="/">
              <div>Manipulate</div>
            </NavLink>
          </li>
          <li className="duration-200 hover:scale-110 hover:underline decoration-strawberry mx-3">
            <NavLink to="/">
              <div>Flips</div>
            </NavLink>
          </li>
          <li className="duration-200 hover:scale-110 hover:underline decoration-strawberry mx-3">
            <NavLink to="/users">
              <div>
                Bazaar<span className="text-strawberry">{" → "}</span>NPC
              </div>
            </NavLink>
          </li>
          <li className="duration-200 hover:scale-110 hover:underline decoration-strawberry mx-3">
            <NavLink to="/users">
              <div>
                NPC<span className="text-strawberry">{" → "}</span>Bazaar
              </div>
            </NavLink>
          </li>
          <li className="duration-200 hover:scale-110 hover:underline decoration-strawberry mx-3">
            <NavLink to="/users">
              <div>Crafts</div>
            </NavLink>
          </li>
          <li className="duration-200 hover:scale-110 hover:underline decoration-strawberry mx-3">
            <NavLink to="/users">
              <div>Crash</div>
            </NavLink>
          </li>
          <li className="duration-200 hover:scale-110 hover:underline decoration-strawberry mx-3">
            <NavLink to="/users">
              <div>AH</div>
            </NavLink>
          </li>
        </ul>
      </div>
      <div className="relative w-fit flex mt-3">
        <button
          onClick={toggleNavbar}
          className="ml-3 animate-button_touch_up text-5xl text-gray duration-500 active:animate-button_touch_down active:fill-mode-forwards"
        >
          <svg
            fill="#ffffff"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            width="24"
            height="24"
          >
            <path fill="none" d="M0 0h24v24H0z" />
            <path d="M14 10v4h-4v-4h4zm2 0h5v4h-5v-4zm-2 11h-4v-5h4v5zm2 0v-5h5v4a1 1 0 0 1-1 1h-4zM14 3v5h-4V3h4zm2 0h4a1 1 0 0 1 1 1v4h-5V3zm-8 7v4H3v-4h5zm0 11H4a1 1 0 0 1-1-1v-4h5v5zM8 3v5H3V4a1 1 0 0 1 1-1h4z" />
          </svg>
        </button>
        <div className="ml-5 mt-0.5 text-white font-groovy font-bold tracking-widest">
          <NavLink to="/">Watchdawg</NavLink>
        </div>
      </div>
    </nav>
  );
}
