import axios from "axios"
import { GetServerSideProps } from "next"
import CardList from "../../components/CardList"
import Layout from "../../components/Layout"
import UserProfile from "../../components/UserProfile"
import { Post } from "../../types/Post"
import { User } from "../../types/User"

type Props = {
  posts: Post[]
  user: User
}

const PostId: React.FC<Props> = ({ posts, user }) => {
  if (!user) return <></>
  return (
    <Layout>
      <div className="max-w-5xl m-auto">
        <div className="flex gap-4">
          <div className="w-1/3">
            <UserProfile user={user} />
          </div>
          <div className="w-2/3">
            <CardList posts={posts} />
          </div>
        </div>
      </div>
    </Layout>
  )
}

export const getServerSideProps: GetServerSideProps = async (context) => {
  if (!context.params) return { props: {} }
  const { userId } = context.params;
  const res = await axios.get(`http://localhost/${userId}/posts`)
  const res2 = await axios.get(`http://localhost/${userId}`)

  return { props: { posts: res.data, user: res2.data } }
}

export default PostId