package client

//enum ClientEventType {
//StatusChanged = 'status-changed',
//DocumentsChanged = 'documents-changed',
//PeersChanged = 'peers-changed',
//StreamConnectionStatusChanged = 'stream-connection-status-changed',
//DocumentSyncResult = 'document-sync-result',
//}
//e

type EventType string

const (
	StatusChanged                 EventType = "status-changed"
	DocumentsChanged              EventType = "documents-changed"
	PeersChanged                  EventType = "peers-changed"
	StreamConnectionStatusChanged EventType = "stream-connection-status-changed"
	DocumentSyncResult            EventType = "document-sync-result"
)
