import Link from "next/link"

type Props = {
  rankingTitle: string
  rankingUnit: string
  contents: RankingContent[]
}

type RankingContent = {
  img: string
  name: string
  count: number
}

const Ranking: React.FC<Props> = ({ rankingTitle, rankingUnit, contents }) => {
  return (
    <div className="w-full">
      <span className="font-bold">{rankingTitle}</span>
      <span>週間</span>
      <span>月間</span>
      <ul className="flex flex-col mt-4 p-2 justify-between">
        {contents.map((content, index) => {
          return (
            <Link href="">
              <a>
                <li className="flex gap-8 hover:bg-gray-200">
                  <span>
                    {index + 1}
                  </span>
                  <span>
                    {content.name}
                  </span>
                  <div>
                    <span className="font-bold">{content.count}</span>
                    <span>{rankingUnit}</span>
                  </div>
                </li>
              </a>
            </Link>
          )
        })}
      </ul>
      <hr />
    </div>
  )
}

export default Ranking
