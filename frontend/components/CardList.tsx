import Card from "./Card"

type Props = {}

const CardList: React.FC<Props> = ({ }) => {
  return (
    <div className="flex flex-col gap-4">
      <Card />
      <Card />
      <Card />
    </div>
  )
}

export default CardList
