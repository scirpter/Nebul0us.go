/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./app/**/*.{ts,tsx,jsx,js}"],
  theme: {
    extend: {
      backgroundImage: {
        twinkle: "url('/assets/images/twinkle.gif')",
      },
      screens: {
        tablet: "640px",
        laptop: "1024px",
        desktop: "1280px",
      },
      colors: {
        dark_purple: "#4A2C4C",
        total_pink: "#FF00FF",
        hot_pink: "#FF0080",
        light_total_pink: "#FFABFF",
        purple: "#4D32AA",
        clean_purple: "#A675FF",
        purple_cloud: "#DED3E4",
        clean_blue: "#515AFC",
        dark_gray: "#000E12",
        midnight_purple: "#040005",
        gray: "#A9A9A9",
        mint: "#00A170",
        dead_mint: "#00DEAD",
        embed_background: "#2F3136",
        light_mint: "#00FFA6",
        honeysucle: "#D65076",
        clean_black: "#0D0D0D",
        light_clean_yellow: "#FFFDA4",
        discord_idle: "#F4A219",
        discord_online: "#3BA45C",
        discord_offline_outer: "#717C89",
        discord_offline_inner: "#393C42",
        lighter_clean_black: "#1A1A1A",
        strawberry: "#FF1F7A",
        light_sundown_orange: "#F9BEAF",
        love_red: "#E41B17",
        better_dark_purple: "#0A090D",
        darkmode_light: "#17191e",
        darkmode_dark: "#0c0d0f",
        darkmode_mid: "#0c0e12",
      },

      brightness: {
        25: ".25",
      },

      keyframes: {
        rbwiggle: {
          "0%": {
            textShadow:
              "0 0 5px #fff, -3px 1px 0 #ca00af, -2.5px 2px 0 #ca00af, -1.5px 3px 0 #ca00af",
            color: "#ffffff",
          },
          "30%": {
            textShadow:
              "0 0 5px #00e1ff, 3px -1px 0 #00e1ff, 2.5px -2px 0 #00e1ff, 1.5px -3px 0 #00e1ff, 2.5px -2px 0 #00e1ff, 3px -1px 0 #00e1ff",
            color: "#d9f6fa",
          },
          "100%": {
            textShadow:
              "0 0 5px #fff, -3px 1px 0 #ca00af, -2.5px 2px 0 #ca00af, -1.5px 3px 0 #ca00af",
            color: "rgb(255, 235, 254)",
          },
        },
        site_match_wiggle: {
          "0%": {
            textShadow: "-3px 1px 0 #bbb, -2.5px 2px 0 #bbb, -1.5px 3px 0 #bbb",
            color: "#ffffff",
          },
          "30%": {
            textShadow:
              "0 0 5px #FF1F7A, 3px -1px 0 #FF1F7A, 2.5px -2px 0 #FF1F7A, 1.5px -3px 0 #FF1F7A, 2.5px -2px 0 #FF1F7A, 3px -1px 0 #FF1F7A",
            color: "#d9f6fa",
          },
          "100%": {
            textShadow: "-3px 1px 0 #bbb, -2.5px 2px 0 #bbb, -1.5px 3px 0 #bbb",
            color: "rgb(255, 235, 254)",
          },
        },
        fade_in: {
          "0%": {
            opacity: "0",
          },
          "100%": {
            opacity: "1",
          },
        },
        button_touch_down: {
          "100%": {
            transform: "rotate(-20deg)",
          },
        },
        button_touch_up: {
          // spin 3-5 times
          "0%": {
            transform: "rotate(-10deg)",
          },
          "100%": {
            transform: "rotate(360deg)",
          },
        },
        error_shake: {
          "0%": {
            transform: "translate(1px, 1px)",
          },
          "10%": {
            transform: "translate(-1px, -2px)",
          },
          "20%": {
            transform: "translate(-3px, 0px)",
          },
          "30%": {
            transform: "translate(3px, 2px)",
          },
          "40%": {
            transform: "translate(1px, -1px)",
          },
          "50%": {
            transform: "translate(-1px, 2px)",
          },
          "60%": {
            transform: "translate(-3px, 1px)",
          },
          "70%": {
            transform: "translate(3px, 1px)",
          },
          "80%": {
            transform: "translate(-1px, -1px)",
          },
          "90%": {
            transform: "translate(1px, 2px)",
          },
          "100%": {
            transform: "translate(1px, -2px)",
          },
        },
      },

      animation: {
        rbwiggle: "rbwiggle 3s ease-in-out infinite",
        site_match_wiggle: "site_match_wiggle 3s ease-in-out infinite",
        fade_in: "fade_in 1s ease-in-out",
        button_touch_down: "button_touch_down 0.3s ease-in-out",
        button_touch_up: "button_touch_up 0.5s ease-in-out",
        error_shake: "error_shake 0.2s ease-in-out",
      },

      fontFamily: {
        ultra: ["Ultra", "monospace"],
        groovy: ["Balancegroovy", "monospace"],
        gta: ["pricedown", "monospace"],
        monospace: ["Roboto Mono", "monospace"],
      },
    },
  },
  plugins: [require("tailwindcss-animate")],
};
