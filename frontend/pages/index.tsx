import type { GetServerSideProps, NextPage } from 'next'
import Image from "next/image"

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
        <div className="w-1/5">
          <Ranking rankingTitle="タグランキング" rankingUnit="users" contents={dummyRanking} />
        </div>
        <div className="w-3/5">
          <CardList posts={posts} />
        </div>
        <div className="w-80">
          <a href="https://lp.nijibox.jp/cp/postdev/" target="_blank">
            <Image className='z-0' src="/adv-3.png" width={300} height={250} />
          </a>
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
