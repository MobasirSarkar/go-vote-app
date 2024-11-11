import ThemeToggler from './components/theme-toggle';
import SignUpForm from './components/sign-up-form';
import { ThemeProvider } from './hooks/ThemeProvier';
import SignInForm from './components/sign-in-form';
function App() {
   return (

      <div className="hero bg-base-200 min-h-screen">
         <div className="hero-content flex-col">
            <ThemeToggler />
            <SignUpForm />
         </div>
      </div>
   )
}
export default App
