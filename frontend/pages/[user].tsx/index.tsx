import { NextPage,GetStaticProps } from "next"
import Layout from "../../components/Layout"

type Props = {}

const User: NextPage<Props> = ({ }) => {
  return (
    <Layout title="user">
    </Layout>
  )
}

export default User

export const getStaticProps: GetStaticProps = async (ctx) => {
  const fuga = "fuga"
  const props = await getOGPStaticProps(Hoge, {
    fuga: fuga,^
  })

  return { props }
}
