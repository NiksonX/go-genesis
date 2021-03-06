package metric

import (
	"strconv"
	"time"

	"github.com/GenesisCommunity/go-genesis/packages/consts"
	"github.com/GenesisCommunity/go-genesis/packages/model"

	log "github.com/sirupsen/logrus"
)

const (
	metricEcosystemPages   = "ecosystem_pages"
	metricEcosystemMembers = "ecosystem_members"
	metricEcosystemTx      = "ecosystem_tx"
)

// CollectMetricDataForEcosystemTables returns metrics for some tables of ecosystems
func CollectMetricDataForEcosystemTables() (metricValues []*Value, err error) {
	stateIDs, err := model.GetAllSystemStatesIDs()
	if err != nil {
		log.WithFields(log.Fields{"error": err, "type": consts.DBError}).Error("get all system states ids")
		return nil, err
	}

	now := time.Now()
	unixDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Unix()

	for _, stateID := range stateIDs {
		var pagesCount, membersCount int64

		tablePrefix := strconv.FormatInt(stateID, 10)

		p := &model.Page{}
		p.SetTablePrefix(tablePrefix)
		if pagesCount, err = p.Count(); err != nil {
			log.WithFields(log.Fields{"error": err, "type": consts.DBError}).Error("get count of pages")
			return nil, err
		}
		metricValues = append(metricValues, &Value{
			Time:   unixDate,
			Metric: metricEcosystemPages,
			Key:    tablePrefix,
			Value:  pagesCount,
		})

		m := &model.Member{}
		m.SetTablePrefix(tablePrefix)
		if membersCount, err = m.Count(); err != nil {
			log.WithFields(log.Fields{"error": err, "type": consts.DBError}).Error("get count of members")
			return nil, err
		}
		metricValues = append(metricValues, &Value{
			Time:   unixDate,
			Metric: metricEcosystemMembers,
			Key:    tablePrefix,
			Value:  membersCount,
		})
	}

	return metricValues, nil
}

// CollectMetricDataForEcosystemTx returns metrics for transactions of ecosystems
func CollectMetricDataForEcosystemTx() (metricValues []*Value, err error) {
	ecosystemTx, err := model.GetEcosystemTxPerDay()
	if err != nil {
		log.WithFields(log.Fields{"error": err, "type": consts.DBError}).Error("get ecosystem transactions by period")
		return nil, err
	}
	for _, item := range ecosystemTx {
		if len(item.Ecosystem) == 0 {
			continue
		}

		metricValues = append(metricValues, &Value{
			Time:   item.UnixTime,
			Metric: metricEcosystemTx,
			Key:    item.Ecosystem,
			Value:  item.Count,
		})
	}

	return metricValues, nil
}
