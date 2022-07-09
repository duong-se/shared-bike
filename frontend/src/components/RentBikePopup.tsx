import { Button } from "./Button"

type RentBikePopupProps = {
  isRented: boolean
  onRent: () => void
  onReturn: () => void
  isDisabledAction: boolean
}

export const RentBikePopUp: React.FC<RentBikePopupProps> = ({ isRented, onRent, onReturn }) => {
  return (
    <div>
      <h1>Bike Name</h1>
      <h5>This bike for rent</h5>
      <div className="container mx-2">
        <ol>
          <li>Click on &ldquo;Rent Bike&rdquo;</li>
          <li>Bicycle lock will unlock automatically</li>
          <li>Adjust saddle height</li>
        </ol>
      </div>
      <div className="w-1/2 text-right">
        {!isRented && <Button onClick={onRent} disabled variant="primary">Rent Bike</Button> }
        {isRented && <Button onClick={onReturn} variant="primary">Return Bike</Button>}
      </div>
    </div>
  )
}
