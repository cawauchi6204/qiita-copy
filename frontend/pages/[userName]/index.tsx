import axios from "axios"
import { GetServerSideProps } from "next"
import CardList from "../../components/CardList"
import Layout from "../../components/Layout"
import { Post } from "../../types/Post"

type Props = {
  posts: Post[]
}

const PostId: React.FC<Props> = ({ posts }) => {
  return (
    <Layout>
      <CardList posts={posts} />
    </Layout>
  )
}

export const getServerSideProps: GetServerSideProps = async (context) => {
  const { userName } = context.query;
  const res = await axios.get(`http://localhost/users/${userName}/posts`)

  return { props: { posts: res.data } }
}

export default PostId
