import axios from "axios"
import { GetServerSideProps } from "next"
import Image from "next/image"
import { Facebook, Heart, MoreHorizontal, Pocket, Twitter } from "react-feather"
import ReactMarkdown from 'react-markdown'
import { useRecoilValue } from "recoil"

import Icon from "../../../components/Icon"
import Layout from "../../../components/Layout"
import { userState } from "../../../contexts/UserContext"
import { Post } from "../../../types/Post"

type Props = {
  post: Post
}

const postId: React.FC<Props> = ({ post }) => {
  const user = useRecoilValue(userState);

  const updateLike = async () => {
    try {
      axios.put("/api/like",
        {
          postId: post.id,
          userId: user?.id
        })
    } catch (e) {
      console.error(e)
    }
  }
  return (
    <Layout>
      <div className="flex gap-4">
        <div className="w-1/12 flex flex-col gap-4 items-end pr-6">
          {user && user.id && (
            <>
              <Heart onClick={updateLike} />
              <Pocket />
            </>
          )}
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
            <ReactMarkdown className="mt-12 markdown" children={post.body} />
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
  const { userId, postId } = context.params
  const res = await axios.get<Post>(`http://localhost/user/${userId}/items/${postId}`)
  return { props: { post: res.data } }
}
