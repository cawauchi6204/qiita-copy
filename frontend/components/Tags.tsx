import Link from "next/link"
import { Tag as TagIcon } from "react-feather"

import useTag from "../hooks/useTag"

type Props = {
  postId: string
}

const Tags: React.FC<Props> = ({ postId }) => {
  const { tags } = useTag(postId)
  return (
    <>
      {
        tags.length > 0 && (
          <div className="mt-4 text-gray-600 flex gap-2 items-center">
            <TagIcon className="inline" size={20} />
            {tags.map((tag) => (<Link key={tag.id} href={`/tags/${tag.tagId}`}><a><span>{tag.tagId},</span></a></Link>))}
          </div>
        )
      }
    </>
  )
}

export default Tags
