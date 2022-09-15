import axios from "axios"
import { GetServerSideProps } from "next"
import Image from "next/image"

import CardList from "../../components/CardList"
import Layout from "../../components/Layout"
import { Post } from "../../types/Post"
import { Tag } from "../../types/Tag"

type Props = {
  tag: Tag
  posts: Post[]
}

const TagId: React.FC<Props> = ({ tag, posts }) => {
  return (
    <Layout className="flex gap-6">
      <div className="w-1/3">
        <div className="flex flex-col bg-white text-center p-6">
          <div className="mx-auto">
            <Image className='z-index-0' src="/adv-1.png" width={150} height={150} />
          </div>
          <p className="mt-2 font-bold">
            {tag.id}
          </p>
          <div className="flex m-auto mt-6 justify-around">
            <div className="flex flex-col w-full">1780<br />記事</div>
            <div className="flex flex-col w-full">657<br />フォロワー</div>
          </div>
          <div className="flex gap-4 m-auto">
            <button className="text-xs text-green-600 border-2 border-green-600 border-solid p-2">記事をフォローする</button>
            <button className="text-xs text-green-600 border-2 border-green-600 border-solid p-2">質問をフォローする</button>
          </div>
        </div>
        <div>
          <p>ランキぐん</p>
        </div>
      </div>
      <div className="w-2/3">
        <CardList posts={posts} />
      </div>
    </Layout>
  )
}

export default TagId

export const getServerSideProps: GetServerSideProps = async (context) => {
  if (!context.params) return { props: {} }
  const { tagId } = context.params
  const tag = await axios.get<Tag>(
    encodeURI(
      `http://localhost/tag/${tagId}`
    ))
  const posts = await axios.get<Post[]>(
    encodeURI(
      `http://localhost/post/tags/${tagId}`
    ))
  return { props: { tag: tag.data, posts: posts.data } }
}
