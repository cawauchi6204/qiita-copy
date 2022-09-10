import axios from "axios"
import { GetServerSideProps } from "next"
import Image from "next/image"
import { Facebook, Heart, MoreHorizontal, Pocket, Twitter } from "react-feather"
import ReactMarkdown from 'react-markdown'

import Icon from "../../../components/Icon"
import Layout from "../../../components/Layout"
import { Post } from "../../../types/Post"

type Props = {
  post: Post
}

const postId: React.FC<Props> = ({ post }) => {
  return (
    <Layout>
      <div className="flex gap-4">
        <div className="w-1/12 flex flex-col gap-4 items-end pr-6">
          <Heart />
          <Pocket />
          <Twitter />
          <Facebook />
          <MoreHorizontal />
        </div>
        <div className="w-8/12">
          <div className="w-full bg-white p-8">
            <div className="flex gap-2">
              <Icon width={20} height={20} />
              <span>{post.postedBy}</span>
            </div>
            <h1 className="text-3xl font-bold">{post.title}</h1>
            <ReactMarkdown className="mt-12" children={post.body} />
          </div>
        </div>
        <div className="w-3/12 flex flex-col gap-16">
          <Image src="/adv-1.png" width={300} height={250} />
          <Image src="/adv-2.png" width={300} height={250} />
        </div>
      </div>
    </Layout>
  )
}

export default postId

export const getServerSideProps: GetServerSideProps = async (context) => {
  if (!context.params) return { props: {} }
  const { userName, postId } = context.params
  const res = await axios.get<Post>(`http://localhost/${userName}/items/${postId}`)
  return { props: { post: res.data } }
}