import Head from "next/head"
import Header from "./Header"
import Footer from "./Footer"

type Props = {
  title?: string
  children: React.ReactNode
}

const Layout: React.FC<Props> = ({ title = "Qiita", children }) => {
  return (
    <>
      <Head>
        <title>{title}</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/tori.png" />
      </Head>
      <Header />
      <main className="px-8 py-6 max-w-7xl m-auto">
        {children}
      </main>
      <Footer />
    </>
  )
}

export default Layout
