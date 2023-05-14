import Footer from "~/components/footer";
import Navbar from "~/components/navbar";
import ParticlesComponent from "~/components/particles";
import { Link } from "@remix-run/react";
import type { V2_MetaFunction } from "@remix-run/node";
import PageBg from "~/components/page_bg";

export const meta: V2_MetaFunction = () => {
  return [
    {
      title: "Home",
    },
  ];
};

export default function Index() {
  return (
    <main className="relative min-h-screen">
      <PageBg />
      <ParticlesComponent />
      <section className="relative animate-fade_in flex h-full justify-center pt-16">
        <div>
          <h1 className="relative font-bold animate-site_match_wiggle cursor-default select-none text-center font-ultra text-4xl text-white tablet:text-5xl laptop:text-7xl">
            Watchdawg
          </h1>
          <p className="mt-14 cursor-default select-none text-center font-ultra font-bold text-lg text-white tablet:text-2xl desktop:text-3xl">
            Renegade Grade Bazaar Artillery
          </p>
          <ul className="mt-16 grid grid-cols-2 justify-between gap-2 text-white bg-darkmode_dark w-3/4 laptop:w-1/2 mx-auto">
            {[{ span: "Tech Stack", to: "/about/stack" }].map((item, index) => (
              <li key={index} className="text-center last:col-span-full">
                <Link to={item.to}>
                  <div className="menu-button group/btn relative overflow-hidden rounded border border-white py-1.5 duration-500 hover:text-black">
                    <div className="absolute -left-72 duration-500 group-hover/btn:scale-[5.0]">
                      <div className="bg-white w-96 h-96 -rotate-45"></div>
                    </div>
                    <div className="button-text relative">
                      <span className="truncate font-groovy tablet:text-lg">
                        {item.span}
                      </span>
                    </div>
                  </div>
                </Link>
              </li>
            ))}
          </ul>
        </div>
      </section>
      <Navbar fixed={true} />
      <Footer />
    </main>
  );
}
