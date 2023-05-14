import * as React from "react";
import Particles from "react-particles";
import type { Engine } from "tsparticles-engine";
import { loadFull } from "tsparticles";

export default function ParticlesComponent() {
  const particlesInit = React.useCallback(async (engine: Engine) => {
    // console.log(engine);

    // you can initialize the tsParticles instance (engine) here, adding custom shapes or presets
    // this loads the tsparticles package bundle, it's the easiest method for getting everything ready
    // starting from v2 you can add only the features you need reducing the bundle size
    await loadFull(engine);
  }, []);

  return (
    <Particles
      id="tsparticles"
      init={particlesInit}
      options={{
        /*
        background: {
          color: {
            value: "#000000",
          },
          opacity: 0,
        },
        */
        fpsLimit: 50,
        interactivity: {
          events: {
            onClick: {
              enable: false,
              mode: "push",
            },
            onHover: {
              enable: true,
              mode: "repulse",
            },
            resize: true,
          },
          modes: {
            push: {
              quantity: 0,
            },
            repulse: {
              distance: 150,
              duration: 0.4,
            },
          },
        },
        particles: {
          color: {
            value: "#ffffff",
          },
          links: {
            color: "#ffffff",
            distance: 150,
            enable: false,
            opacity: 0.5,
            width: 1, // default: 1
          },
          collisions: {
            enable: true,
          },
          move: {
            direction: "none",
            enable: true,
            outModes: {
              default: "bounce",
            },
            random: false,
            speed: 0.2,
            straight: false,
          },
          number: {
            density: {
              enable: true,
              area: 800,
            },
            value: 120,
          },
          opacity: {
            value: { min: 0.5, max: 0.8 },
          },
          interactivity: {
            modes: {
              bubble: {
                distance: 400,
                duration: 2,
                opacity: 0.8,
                size: 40,
              },
              grab: {
                distance: 400,
                links: {
                  opacity: 1,
                },
              },
            },
          },
          shape: {
            type: "circle",
          },
          size: {
            value: { min: 0.2, max: 1.69 },
          },
        },
        detectRetina: true,
      }}
    />
  );
}
