import apis, { Isu, GraphRequest } from '../../lib/apis'
import { useCallback } from 'react'
import NowLoading from '../UI/NowLoading'
import TransitionGraph from './TransitionGraph'
import SittingGraph from './SittingGraph'
import Score from './Score'
import DateInput from './DateInput'
import useGraph from './use/graph'

interface Props {
  isu: Isu
}

const IsuGraphCardContent = ({ isu }: Props) => {
  const getGraphs = useCallback(
    (params: GraphRequest) => {
      return apis.getIsuGraphs(isu.jia_isu_uuid, params)
    },
    [isu.jia_isu_uuid]
  )

  const {
    graphs,
    transitionData,
    sittingData,
    timeCategories,
    score,
    day,
    tooltipData,
    fetchGraphs
  } = useGraph(getGraphs)

  if (graphs.length === 0) {
    return <NowLoading />
  }
  return (
    <div className="flex flex-col gap-12">
      <div className="flex">
        <DateInput day={day} fetchGraphs={fetchGraphs} />
      </div>
      <div className="flex flex-col gap-8">
        <TransitionGraph
          transitionData={transitionData}
          timeCategories={timeCategories}
          tooltipData={tooltipData}
        />
        <SittingGraph
          sittingData={sittingData}
          timeCategories={timeCategories}
        />
        <Score score={score} />
      </div>
    </div>
  )
}

export default IsuGraphCardContent
