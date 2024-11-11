import { createContext, useContext, useEffect, useState } from "react"

type Theme = "dark" | "cupcake" // Restrict themes to dark and cupcake only

type ThemeProviderProps = {
   children: React.ReactNode
   defaultTheme?: Theme
   storageKey?: string
}

type ThemeProviderState = {
   theme: Theme
   setTheme: (theme: Theme) => void
}

const initialState: ThemeProviderState = {
   theme: "dark", // Default theme is "cupcake"
   setTheme: () => null,
}

const ThemeProviderContext = createContext<ThemeProviderState>(initialState)

export function ThemeProvider({
   children,
   defaultTheme = "dark", // Default to "cupcake" theme
   storageKey = "vite-ui-theme",
   ...props
}: ThemeProviderProps) {
   const [theme, setTheme] = useState<Theme>(
      () => (localStorage.getItem(storageKey) as Theme) || defaultTheme
   )

   useEffect(() => {
      const root = window.document.documentElement

      root.classList.remove("cupcake", "dark") // Remove both cupcake and dark themes

      // Apply the selected theme
      root.classList.add(theme)

      // If the theme is changed, store it in localStorage
      localStorage.setItem(storageKey, theme)
   }, [theme])

   const value = {
      theme,
      setTheme: (theme: Theme) => {
         // Update theme in the state and localStorage
         localStorage.setItem(storageKey, theme)
         setTheme(theme)
      },
   }

   return (
      <ThemeProviderContext.Provider {...props} value={value}>
         {children}
      </ThemeProviderContext.Provider>
   )
}

export const useTheme = () => {
   const context = useContext(ThemeProviderContext)

   if (context === undefined)
      throw new Error("useTheme must be used within a ThemeProvider")

   return context
}
