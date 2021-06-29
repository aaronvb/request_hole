module.exports = {
  purge: ['./src/**/*.{js,jsx,ts,tsx}', './public/index.html'],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {
      keyframes: {
        'slide-right': {
          '0%': {
            transform: 'translateX(-10px)'
          },
          '100%': {
            transform: 'translateX(0)'
          },
        }
      },
      animation: {
        'slide-right': 'slide-right 0.5s ease-out'
      }
    },
  },
  variants: {
    extend: {},
  },
  plugins: [],
}
