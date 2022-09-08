import { GetServerSideProps } from "next"

type Props = {}

// ここでuser_idとpost_idを全て持ってくる処理をする
// SSGだとビルド時間がエグくなりそう
// SSRにする
// paramからユーザーのIDを取得
// getAllPostsByUserID
// ユーザーIDからそのユーザーが書いた記事を取得してくるAPIを叩く

const PostId: React.FC<Props> = ({ }) => {
  return (
    <>
    </>
  )
}

export const getServerSideProps: GetServerSideProps = async (context) => {
  // Fetch data from external API
  const res = await axios.get(`https://localhost/`)
  const data = await res.json()

  // Pass data to the page via props
  return { props: { data } }
}

export default PostId
