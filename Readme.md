# godipInfluence

A package that allows the calculation of influenced provinces of a [zond/godip](https://github.com/zond/godip) game. I decided to not introduce this into the core as it is only a matter of how a user would like to see it implemented not of the adjudication. Feel free to discuss in the issues.

## Installation

```bash
go get github.com/wulfheart/godip-influence
```

## Usage

```golang
import (
	"github.com/wulfheart/godip-influence/defaultInfluences"
	"github.com/wulfheart/godip-influence/influenceCalculators"
	"github.com/zond/godip/variants"
)

// Variant/game starting
variantName := "Classical"
variant := variants.Variants[variantName]
state, _ := variant.Start()

// Getting the influence
// There may be custom functions of type InfluenceCalculator func (old Influence, new *state.State) Influence
var influence = influenceCalculators.WebdiplomacyClassic(defaultInfluences.ConvertToInfluence(defaultInfluences.Classical), state)
// e.g. influence["par"] -> godip.France
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[Licensed under the EUPL-1.2-or-later](https://choosealicense.com/licenses/eupl-1.2/)