import type { NextPage } from 'next'
import Layout from "../components/Layout"
import CardList from "../components/CardList"

const Home: NextPage = () => {
  return (
    <Layout>
      <main>
        <CardList />
      </main>
    </Layout>
  )
}

export default Home
