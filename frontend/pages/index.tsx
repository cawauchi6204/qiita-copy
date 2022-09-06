import type { NextPage } from 'next'
import Layout from "../components/Layout"

const Home: NextPage = () => {
  return (
    <Layout>
      <main>
        <button className="bg-green-600 hover:bg-green-500 text-white rounded px-4 py-2">登録する</button>
      </main>
    </Layout>
  )
}

export default Home
