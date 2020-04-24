package downloader

import (
	"log"

	"git.vsh-labs.cz/jelito/sarinka/src/storage"
	"git.vsh-labs.cz/jelito/sarinka/src/youtube"
)

func getChannelIds() []string {
	return []string{
		//"UCV9FxSrPVEX3vjBax04V_uA", // booba
		//"UCvMF0dbLxjQVg9xCxIx8E6g", // learn english with om nom
		//"UCdknlZNiH-2VWrvjXkx9dJA", // pocoyo english
		//"UCbCmjCuTUZos6Inko4u57UQ", //coco melon
		//"UC-Gm4EN7nNNR3k67J8ywF4g", // blippi toys
		//"UC5PYHgAzJ1wLEidB58SK6Xw", // blippi
		//"UCH5TSOM_0xZbbLImGqigtsA", // vroom vroom
		//"UCgNqMjgfN6YQYCYOxhp78xQ", // pisnicky pro deti
		//"UCyTOz12LseUJhDtl8K6-vFw", // baby studio
		//"UCnUExytdcrLyl_ccqe2dxpg", // ceske pohadky
		//"UCVGHomAxlCzmBH7T5SN0_ag", // ciperkove
		//"UC20HC7Rj2_pyyUq4xAjtw0Q", // baby smile
		//"UCRdYoBTTupd3svI9kV36UvQ", // nakladak
		//"UCjuM-aDNfzbVShImGfj0WOA", // odtahove auto
		//"UClroLesWYk7cc9Tctv8HrkA", // monster auto
		//"UCZBqWU1GgUHBTPZ-FMPZ8-Q", // byl jednou
		//"UC6zhI71atP7YLoZyIyCIGNw", // dave and ava
		//"UCwCP0VQHMl0YvZJCAUB6pPg", // cat family

		//"UCEzJkg_EtsuPwDnez851ZKw", // pocoyo songs - nic tam neni
	}
}

func getPlaylistIds() []string {
	return []string{
		//"PL4nt7KiBPAY-CA0Ll2boMp4wM--Lb11XC", // pocoyo songs
		// Nefunguje //"PLql8Ul6KKObLH2fYr6WxuKQD8q2a176Hd", // krtek
		//"PL7u4k6y5_Re-LBijWl2V_FJeMlvLPO9Wg", // fik
		//"PLxf0FUnNFREf3zhBkp9vhpAvAWL8IWvIM", // donald
		//"PL_c00IPCdRoGD0k6aG7725OaqHQQffcVf", // byl jednou jeden zivot
		//"PL74Iv3ZxWVVmhDuhyIETnbEeyNWhETfXR", // byl jednou jeden objevitel
	}
}

type Handler struct {
	storage *storage.Handler
	youtube *youtube.Handler
}

func New(storage *storage.Handler, youtube *youtube.Handler) *Handler {
	return &Handler{
		storage: storage,
		youtube: youtube,
	}
}

func (h *Handler) SaveData() {
	for _, channelId := range getChannelIds() {
		err := h.saveChannelData(channelId)
		if err != nil {
			log.Println(err)
		}
	}
	for _, playlistId := range getPlaylistIds() {
		err := h.savePlaylistData(playlistId)
		if err != nil {
			log.Println(err)
		}
	}
}

func (h *Handler) saveChannelData(channelId string) error {
	res, err := h.youtube.LoadChannelData(channelId)
	if err != nil {
		return err
	}

	return h.storage.SaveData(channelId, res)
}

func (h *Handler) savePlaylistData(playlistId string) error {
	res, err := h.youtube.LoadPlaylistData(playlistId)
	if err != nil {
		return err
	}

	return h.storage.SaveData(playlistId, res)
}
