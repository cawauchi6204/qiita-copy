import type { GetServerSideProps, NextPage } from 'next'
import Layout from "../components/Layout"
import CardList from "../components/CardList"
import Ranking from "../components/Ranking"
import axios from 'axios'
import { Post } from '../types/Post'

const dummyRanking = [
  {
    img: "",
    name: "Python",
    count: 127,
  },
  {
    img: "",
    name: "Python",
    count: 127,
  },
  {
    img: "",
    name: "Python",
    count: 127,
  },
  {
    img: "",
    name: "Python",
    count: 127,
  },
  {
    img: "",
    name: "Python",
    count: 127,
  },
]

type Props = {
  posts: Post[]
}

const Home: NextPage<Props> = ({ posts }) => {
  return (
    <Layout>
      <div className="flex gap-8">
        <div className="w-1/6">
          <Ranking rankingTitle="タグランキング" rankingUnit="users" contents={dummyRanking} />
        </div>
        <div className="w-2/3">
          <CardList posts={posts} />
        </div>
        <div className="w-1/6">
          <Ranking rankingTitle="タグランキング" rankingUnit="users" contents={dummyRanking} />
        </div>
      </div>
    </Layout>
  )
}

export const getServerSideProps: GetServerSideProps = async (context) => {
  const res = await axios.get(`http://localhost/posts`)
  return { props: { posts: res.data } }
}

export default Home
