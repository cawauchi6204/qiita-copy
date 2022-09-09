import '../styles/globals.css'
import type { AppProps } from 'next/app'
import {
  RecoilRoot,
} from 'recoil';
import UserContext from '../contexts/UserContext';

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <RecoilRoot>
      <UserContext />
      <Component {...pageProps} />
    </RecoilRoot>
  )
}
export default MyApp
