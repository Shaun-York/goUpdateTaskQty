package gqldocs

// MachineByPk MachineByPk
type MachineByPk struct {
	 Machine `json:"Machines_by_pk"`
}
//UpdatedTaskByPk UpdatedTaskByPk
type UpdatedTaskByPk struct {
	 Task `json:"update_Tasks_by_pk"`
}
// Machine machine
type Machine struct {
	InternalID string `json:"internalid"`
	Location string `json:"location_display"`
	Name string `json:"name"`
	Workcenter string `json:"workcenter_display"`
	Tasks []Task `json:"MachineTasks"`
}
// Task task
type Task struct {
	InternalID 			string `json:"internalid"`
	Priority 			string `json:"priority"`
	Machine 			string `json:"machine"`
	MachineDisplay 		string `json:"machine_display"`
	Location 			string `json:"location"`
	LocationDisplay 	string `json:"location_display"`
	Workcenter 			string `json:"workcenter"`
	WorkcenterDisplay 	string `json:"workcenter_display"`
	Workorder 			string `json:"workorder"`
	WorkorderDisplay 	string `json:"workorder_display"`
	Item 				string `json:"item"`
	ItemDisplay 		string `json:"item_display"`
	Task 				string `json:"task"`
	Status 				string `json:"status"`
	StatusColor 		string `json:"state_color"`
	Sequence 			string `json:"sequence"`
	SetupTime 			string `json:"setupTime"`
	LaborTime 			string `json:"laborTime"`
	InputQty 			string `json:"inputQty"`
	CompletedQty 		string `json:"completedQty"`
	ArtURL 				string `json:"art_url"`
	Ped 				string `json:"ped"`
}

// AddWocID update completion row with workordercompletion id
// args:
//id: - completion pk - Int!
//workordercompletion_id: - workordercompletion_id in sqs message - String!
var AddWocID string = `
        mutation ($id: Int!, $workordercompletion_id: String!) {
            update_Completions_by_pk(pk_columns: { id: $id },
                _set: { workordercompletion_id: $workordercompletion_id }
            ) {
                workordercompletion_id
                operation_sequence
                mfgoptask_id
                workorder_id
                completedQty
                location_id
                operator_id
                worktime_id
                machine_id
                workcenter
                item_id
                action
                id
            }
        }
	`;
// DeleteCompletion delete completion
var DeleteCompletion string = `
			mutation remove($id: Int!){delete_Completions_by_pk(id: $id) {
                    workordercompletion_id
                    operation_sequence
                    mfgoptask_id
                    workorder_id
                    completedQty
                    location_id
                    operator_id
                    worktime_id
                    machine_id
                    workcenter
                    item_id
                    action
                    id
                    }
                }
			`;
//LastCompletion find last woc		
var LastCompletion string = `
        query lastCompletion($workorder_id: String!, $operation_sequence: String!, $operator_id: String!, $action: String!) {
            Completions(limit: 1 order_by: { id: desc_nulls_first },
            where: {
                    workordercompletion_id: { _neq: "" },
                    operation_sequence: { _eq: $operation_sequence },
    	            workorder_id: { _eq: $workorder_id },
                    operator_id: { _eq: $operator_id },
                    action: { _eq: $action }
  	            }){
                      workordercompletion_id
                      completedQty
                      id
                  }
                }
	`;
// GetCompletedQty get completed quanity
var GetCompletedQty string = `
        query completedQty($internalid: String!) {
            Tasks_by_pk(internalid: $internalid) {
                completedQty
                inputQty
            }
        }
	`;
// GetMachine props
var GetMachine string = `
    query ($id: String!) { 
		Machines_by_pk(internalid: $id) {
		internalid
		name
		workcenter_display
		location_display
    	MachineTasks (limit: 1 order_by: { ped: asc workorder_display: desc }) {
      		internalid
      		priority
      		machine
      		machine_display
      		location
      		location_display
      		workcenter
      		workcenter_display
      		workorder
      		workorder_display
      		item
      		item_display
      		task
      		status
      		state_color
      		sequence
      		setupTime
      		laborTime
      		inputQty
      		completedQty
      		art_url
      		ped
    	}
	  }
	}
`;
// UpdateOperationCompletedQty update Tasks table row completedQty with pk
var UpdateOperationCompletedQty string = `
        mutation ($internalid: String!, $completedQty: String!){ 
            update_Tasks_by_pk(
                pk_columns: {internalid: $internalid}, 
                _set: {completedQty: $completedQty}){
                    internalid
                    completedQty
                }
		}`
// DeletedCompleted delete row in Tasks table by pk
var DeletedCompleted string = `
        mutation delTask($internalid: String!){ 
            delete_Tasks_by_pk(internalid: $internalid) {
                internalid
                workorder
                workcenter
                location
            }
		}`
