/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./pages/*.templ", "./pages/*.html"],
  theme: {
    fontFamily: {
        ArchivoSemiCondensed_Regular: ['"Archivo SemiCondensed-regular"'],
        ArchivoSemiCondensed_SemiBold: ['"Archivo SemiCondensed-semibold"'],
        ArchivoSemiCondensed_Thin: ['"Archivo SemiCondensed-thin"']
      },
      extend: {
        width: {
      },
      fontSize: {
      },
      gridTemplateColumns: {
        'auto-fill-minmax': 'repeat(auto-fill, minmax(337px, 1fr))',
        'bodyLayout': '18.875rem 1fr',
        'mintExtensionPage': '2fr 1.5fr',
      },
      spacing: {
      },
       fontFamily: {
      },
       colors: {
      },
      backgroundImage: {
        'CreateToken22': 'linear-gradient(88.93deg, #9945FF -4.48%, #12DB88 114.18%)',
        'swiperGradient': 'linear-gradient(90deg, #AD6AFF 0%, #12DB88 100%)',
      },
    },
  },
  plugins: [],
}


