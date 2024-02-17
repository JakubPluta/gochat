import '../styles/globals.css'
import { AppProps } from "next/app";
import AuthContextProvider from '@/modules/auth';

export default function App({ Component, pageProps }: AppProps) {
    return (
        <>
        <AuthContextProvider>
        <div className='flex flex-col md:flex-row h-full min-h-screen font-sans'>
            <Component {...pageProps} />
        </div>
        </AuthContextProvider>

    </>
    )
}