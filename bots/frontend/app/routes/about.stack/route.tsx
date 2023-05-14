import Footer from "~/components/footer";
import Navbar from "~/components/navbar";
import ParticlesComponent from "~/components/particles";
import type { V2_MetaFunction } from "@remix-run/node";
import PageBg from "~/components/page_bg";
import ToggleableBox from "~/components/toggleable_box";

export const meta: V2_MetaFunction = () => {
  return [
    {
      title: "Tech Stack",
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
            Tech Stack
          </h1>
          <p className="mt-14 cursor-default select-none text-center font-ultra font-bold text-lg text-white tablet:text-2xl desktop:text-3xl">
            Technologies used in this project
          </p>
          <ul className="mt-16 flex-row space-y-4 text-white bg-darkmode_dark w-3/4 mx-auto mb-32">
            {[
              {
                title: "Golang (backend)",
                description: `Building high-performance, concurrent systems is the only purpose of the programming language known as Go. Building scalable web apps with it is a breeze thanks to its straightforward syntax and integrated concurrency support.
                Over other languages, Go has a number of benefits for a backend. Speed is among the main benefits. Go is a compiled language, which enables it to run programs significantly quicker than interpreted languages like Python or JavaScript.
                For developing high-performance apps that must manage a lot of traffic, this makes it the perfect choice.`,
                image_url: "/assets/images/common/go.png",
              },
              {
                title: "Tailwind CSS (styling)",
                description: `As a utility-first CSS framework, Tailwind CSS offers a wide range of pre-defined CSS classes to help designers build responsive and adaptable UI components.
                  It cuts down on the amount of time and work required for development by enabling developers to quickly prototype and construct user interfaces.`,
                image_url: "/assets/images/common/tailwind.png",
              },
              {
                title: "Remix (frontend)",
                description: `Remix is a framework that provides a powerful and modern tech stack for building web applications.
                It uses a combination of React, TypeScript, and CSS-in-JS to create fast and maintainable UI components.`,
                image_url: "/assets/images/common/remix.png",
              },
            ].map((item, index) => (
              <li key={index} className="text-center last:col-span-full">
                <ToggleableBox
                  title={item.title}
                  imageUrl={item.image_url}
                  defaultState="open"
                >
                  {item.description.split("\n").map((line, index) => (
                    <div key={index}>
                      <p>{line}</p>
                      <br />
                    </div>
                  ))}
                </ToggleableBox>
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
